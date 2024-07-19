'use client'

import NavigationHeader from '@/components/navigation-header'
import { Page } from '@/const/navigation'
import { useEffect, useState } from 'react'
import { searchUsers, User } from './actions'
import SearchInput from './search-input'
import SearchResult from './search-result'

export default function Search() {
  const [searchResult, setSearchResult] = useState<User[]>([])
  const [searchQuery, setSearchQuery] = useState('')

  const search = async (q: string) => {
    const r = await searchUsers(q)
    setSearchResult(r)
  }

  const updateResult = async () => {
    const r = await searchUsers(searchQuery)
    setSearchResult(r)
  }

  useEffect(() => { search(searchQuery) }, [searchQuery])

  return (
    <div className='grid grid-cols-12 bg-black'>
      <div className="col-start-4 col-span-6">
        <NavigationHeader page={Page.Search} />
      </div>
      <div className='col-start-4 col-span-6 py-4'>
        <SearchInput setSearchQuery={setSearchQuery} />
        <SearchResult updateResult={updateResult} searchResult={searchResult} />
      </div>
    </div>
  )
}
