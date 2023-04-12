import { Head } from '@inertiajs/react'

export default function Index({ title }) {
  return (
    <div>
      <Head title="Home" />
      <h1 className="font-bold text-3xl">{title}</h1>
    </div>
  )
}