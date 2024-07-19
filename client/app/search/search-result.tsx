import FollowButton from '@/components/follow-button'
import UnfollowButton from '@/components/unfollow-button'
import UsernameAndID from '@/components/username-and-id'
import { followUser, unfollowUser, User } from './actions'

type Props = {
  updateResult: () => Promise<void>
  searchResult: User[]
}

export default function SearchResult({ updateResult, searchResult }: Props) {
  const follow = async (accountId: string) => {
    await followUser(accountId)
    await updateResult()
  }

  const unfollow = async (accountId: string) => {
    await unfollowUser(accountId)
    await updateResult()
  }

  return (
    <>
      {searchResult.map(({ accountId, username, isFollowed }) => (
        <div key={accountId} className='grid grid-cols-subgrid grid-cols-5 p-4 border-b border-gray-200'>
          <div className='col-span-2 content-center'>
            <UsernameAndID username={username} accountId={accountId} />
          </div>
          <div className='col-start-5 content-center text-right'>
            {
              !isFollowed
                ? <FollowButton onClick={async () => { await follow(accountId) }} />
                : <UnfollowButton onClick={async () => { await unfollow(accountId) }} />
            }
          </div>
        </div>
      ))}
    </>
  )
}
