import { createLazyFileRoute } from '@tanstack/react-router'
import toast from 'react-hot-toast'
import { post } from '../api/request'
import { AIRequest, AIResponse } from '../api/api'
import { useMutation } from '@tanstack/react-query'

export const Route = createLazyFileRoute('/')({
  component: RouteComponent,
})

function RouteComponent() {
  // const [messages, setMessages] = useState<string[]>([])

  const { mutate, error } = useMutation({
    async mutationFn() {
      const resp = await post(
        '/api/ai/prompt',
        { prompt: 'Hello' },
        (obj) => AIRequest.encode(obj).finish(),
        (b) => AIResponse.decode(b)
      )
      console.log(resp)
    },
  })

  return (
    <div className="flex flex-col gap-4">
      <form
        onSubmit={(e) => {
          e.preventDefault()
          const i = e.currentTarget.querySelector('input')
          if (!i) return toast.error('Input not found')
          mutate()
        }}
      >
        <input placeholder="name"></input>
        <button>Create</button>
      </form>
      {error && <div>{error.message}</div>}
    </div>
  )
}
