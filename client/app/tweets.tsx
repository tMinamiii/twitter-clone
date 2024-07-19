import UsernameAndID from '@/components/username-and-id'
import { format } from 'date-fns'
import { Post } from './actions'

type Props = {
  tweets: Post[]
}

export default function Tweets({ tweets }: Props) {
  const formatDate = (date: string): string => {
    const d = new Date(date)
    return format(d, 'yyyy年MM月dd日')
  }
  return (
    <>
      {tweets.map(({ uuid, accountId, username, content, postedAt }) => (
        <div key={uuid} className='p-4 border-b border-gray-200'>
          <div>
            <UsernameAndID username={username} accountId={accountId} />
            <span className='text-gray-400 text-sm pl-2'>{formatDate(postedAt)}</span>
          </div>
          <div className='whitespace-pre-line'>{content}</div>
        </div>
      ))}
    </>
  )
}
