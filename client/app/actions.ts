'use server'

import { getDummySession } from '@/components/actions'
export type Tweets = {
  count: number
  lastUuid: string
  posts: Post[]
}

export type Post = {
  uuid: string;
  accountId: string
  username: string
  content: string
  postedAt: string
}

export const postTweet = async (content: string) => {
  if (content === '') {
    return
  }

  const sessionValue = await getDummySession()

  const body = { content }
  await fetch('http://localhost:1323/v1/posts', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'X-API-Key': String(process.env.API_KEY),
      Cookie: `${process.env.SESSION_NAME}=${sessionValue};`
    },
    body: JSON.stringify(body),
  })
}

export const fetchTweets = async (sinceUuid: string | null): Promise<Tweets> => {
  const q = new URLSearchParams()
  q.set('limit', '20')
  if (sinceUuid) {
    q.set('sinceUuid', sinceUuid)
  }

  const sessionValue = await getDummySession()

  const res = await fetch(`http://localhost:1323/v1/posts/timeline?${q}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'X-API-Key': String(process.env.API_KEY),
      Cookie: `${process.env.SESSION_NAME}=${sessionValue};`
    }
  })

  if (!res.ok) { return { count: 0, lastUuid: '', posts: [] } }

  const resJson = await res.json()
  return {
    count: resJson.count,
    lastUuid: resJson.lastUuid,
    posts: resJson.posts
  }
}
