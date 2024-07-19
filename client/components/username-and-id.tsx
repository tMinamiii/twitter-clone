type Props = {
  username: string
  accountId: string
}

export default function UsernameAndID({ username, accountId }: Props) {
  return (
    <>
      <span className='font-bold'>{username}</span>
      <span className='text-gray-400 text-sm pl-2'>@{accountId}</span>
    </>
  )
}
