type Props = {
  onClick: () => Promise<void>
}

export default function PostButton({ onClick }: Props) {
  return (
    <button
      className='bg-blue-400 hover:bg-blue-800 text-white font-bold py-2 px-4 rounded-full'
      type='button'
      onClick={onClick}
    >
      Post
    </button>
  )
}
