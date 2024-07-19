'use server'

import { getDummySession } from '@/components/actions'

export type User = {
  accountId: string
  username: string
  isFollowed: boolean
}

export const searchUsers = async (username: string) => {
  const params = { username }
  const q = new URLSearchParams(params)

  const sessionValue = await getDummySession()

  const res = await fetch(`http://localhost:1323/v1/users/search?${q}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'X-API-Key': String(process.env.API_KEY),
      Cookie: `${process.env.SESSION_NAME}=${sessionValue};`
    },
  })

  if (!res.ok) return []

  const resJson = await res.json()
  if (resJson.users) {
    const users = resJson.users.map((user: User) => {
      return {
        username: user.username,
        accountId: user.accountId,
        isFollowed: user.isFollowed
      }
    })
    return users
  }
  return []
}

export const followUser = async (accountId: string) => {
  const body = { accountId }

  const sessionValue = await getDummySession()

  await fetch('http://localhost:1323/v1/follows', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'X-API-Key': String(process.env.API_KEY),
      Cookie: `${process.env.SESSION_NAME}=${sessionValue};`
    },
    body: JSON.stringify(body),
  })
}

export const unfollowUser = async (accountId: string) => {
  const params = { accountId }
  const q = new URLSearchParams(params)

  const sessionValue = await getDummySession()

  await fetch(`http://localhost:1323/v1/follows?${q}`, {
    method: 'Delete',
    headers: {
      'Content-Type': 'application/json',
      'X-API-Key': String(process.env.API_KEY),
      Cookie: `${process.env.SESSION_NAME}=${sessionValue};`
    },
  })
}
