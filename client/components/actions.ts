'use server'

import { parseSetCookie } from 'next/dist/compiled/@edge-runtime/cookies'
import { cookies } from 'next/headers'

export const getDummySession = async () => {
  const sessionName = process.env.SESSION_NAME as string
  const sessionValue = cookies().get(sessionName)?.value ?? ''
  if (sessionValue !== '') {
    return sessionValue
  }

  const resp = await fetch('http://localhost:1323/v1/dummy-session', {
    method: 'GET',
    headers: {
      'X-API-Key': String(process.env.API_KEY),
    },
  })

  // サーバーから取得したセッション情報をクッキーに焼く
  const cookieStrings = resp.headers.getSetCookie()
  var dummySession = ''
  for (const cookieString of cookieStrings) {
    const parsed = parseSetCookie(cookieString)
    if (parsed) {
      cookies().set(parsed.name, parsed.value, parsed)
      if (parsed.name == sessionName) {
        dummySession = parsed.value
      }
    }
  }

  return dummySession
}
