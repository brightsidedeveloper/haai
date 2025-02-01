import { createLazyFileRoute } from '@tanstack/react-router'
import { useEffect, useState } from 'react'
import { GetRooms, SocketMessageType, WSMessage } from '../api/api'
import { useQuery } from '@tanstack/react-query'
import { get } from '../api/request'
import toast from 'react-hot-toast'

export const Route = createLazyFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  const { data: rooms, refetch } = useQuery({
    queryKey: ['rooms'],
    queryFn() {
      return get('/api/game/rooms', undefined, (b) => GetRooms.decode(b))
    },
  })

  const socket = useSocket()
  return (
    <div className="flex flex-col gap-4">
      <form
        onSubmit={(e) => {
          e.preventDefault()
          const i = e.currentTarget.querySelector('input')
          if (!i) return toast.error('Input not found')
          const name = i.value
          if (!name) {
            return toast.error('Name is required')
          }

          const createRoom = WSMessage.create({
            type: SocketMessageType.CREATE_ROOM,
            createRoom: { name },
          })

          const encoded = WSMessage.encode(createRoom).finish()

          socket?.send(encoded)
          i.value = ''
          refetch()
        }}
      >
        <input placeholder="name"></input>
        <button>Create</button>
      </form>
      <br />
      <br />
      {rooms?.rooms.map((room) => (
        <div
          key={room.id}
          onClick={() => {
            const joinRoom = WSMessage.create({
              type: SocketMessageType.JOIN_ROOM,
              joinRoom: { id: room.id },
            })

            const encoded = WSMessage.encode(joinRoom).finish()

            socket?.send(encoded)
          }}
        >
          {room.name}
        </div>
      ))}
    </div>
  )
}

function useSocket() {
  const [socket, setSocket] = useState<WebSocket | null>(null)
  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8081/ws')
    socket.onopen = () => {
      console.log('connected')
    }
    socket.onmessage = (e) => {
      const data = new Uint8Array(e.data)

      try {
        const message = WSMessage.decode(data)
        console.log(message)

        switch (message.type) {
          case SocketMessageType.SET_NAME:
            console.log('Received SetNameResponse:', message.setName)
            break

          case SocketMessageType.CREATE_ROOM:
            console.log('Received CreateRoomResponse:', message.createRoom)
            break

          case SocketMessageType.ERROR:
            console.log('WTFFFFFF', message)
            toast.error(message.error?.message ?? 'Unknown error')
            break

          default:
            console.warn('Unknown message type received')
        }
      } catch (err) {
        console.error('Failed to decode message:', err)
      }
    }
    socket.onclose = () => {
      console.log('closed')
    }

    setSocket(socket)

    return () => {
      socket.close()
    }
  }, [])

  return socket
}
