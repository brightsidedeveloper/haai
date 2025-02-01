import { QueryClient } from '@tanstack/react-query'
import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'
import { Toaster } from 'react-hot-toast'

export const Route = createRootRouteWithContext<QueryContext>()({
  component: RootComponent,
})

function RootComponent() {
  return (
    <>
      <Outlet />
      <Toaster />
    </>
  )
}

export type QueryContext = {
  queryClient: QueryClient
}
