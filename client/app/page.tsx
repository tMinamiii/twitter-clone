'use client'

import NavigationHeader from '@/components/navigation-header'
import { Page } from '@/const/navigation'
import { useEffect, useState } from 'react'
import InfiniteScroll from 'react-infinite-scroller'
import { fetchTweets, Post } from './actions'
import PostForm from './post-form'
import Tweets from './tweets'

export default function Timeline() {
  const [lastUuid, setLastUuid] = useState<string | null>(null)
  const [hasMore, setHasMore] = useState<boolean>(true)
  const [tweets, setTweets] = useState<Post[]>([])

  const updateTweets = async () => {
    const tl = await fetchTweets('')
    setLastUuid(tl.lastUuid)
    setTweets(tl.posts)
  }

  const scrollTweets = async () => {
    const tl = await fetchTweets(lastUuid)
    if (tl.posts.length == 0) {
      setHasMore(false)
      return
    }

    setLastUuid(tl.lastUuid)
    setTweets(tweets.concat(tl.posts))
  }

  useEffect(() => { updateTweets() }, [])

  return (
    <div className='grid grid-cols-12 bg-black'>
      <div className="col-start-4 col-span-6">
        <NavigationHeader page={Page.Timeline} />
      </div>
      <div className='col-start-4 col-span-6 py-4'>
        <PostForm updateTimeline={updateTweets} />
        <InfiniteScroll pageStart={0} loadMore={scrollTweets} hasMore={hasMore}>
          <Tweets tweets={tweets} />
        </InfiniteScroll>
      </div>
    </div >
  )
}
