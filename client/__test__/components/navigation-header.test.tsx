import NavigationHeader from '@/components/navigation-header'
import { Page } from '@/const/navigation'
import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react'

describe('NavigationHeader', () => {
  it('ナビゲーションヘッダーが表示され、Timelineの下にバーが表示される', () => {
    render(<NavigationHeader page={Page.Timeline} />)

    const timeline = screen.getByRole('link', { name: 'Timeline' })
    expect(timeline).toBeInTheDocument()
    expect(timeline.classList.toString()).toEqual('col-start-4 text-lg content-center text-center hover:bg-gray-800')
    expect(timeline.getElementsByTagName('span')[0].classList.toString()).toEqual('border-b-4 border-blue-500 py-6')

    const search = screen.getByRole('link', { name: 'Search' })
    expect(search).toBeInTheDocument()
    expect(search.classList.toString()).toEqual('col-start-5 text-lg content-center text-center hover:bg-gray-800')
    expect(search.getElementsByTagName('span')[0].classList.toString()).toEqual('border-b-0')
  })
  it('ナビゲーションヘッダーが表示され、Searchの下にバーが表示される', () => {
    render(<NavigationHeader page={Page.Search} />)

    const timeline = screen.getByRole('link', { name: 'Timeline' })
    expect(timeline).toBeInTheDocument()
    expect(timeline.classList.toString()).toEqual('col-start-4 text-lg content-center text-center hover:bg-gray-800')
    expect(timeline.getElementsByTagName('span')[0].classList.toString()).toEqual('border-b-0')

    const search = screen.getByRole('link', { name: 'Search' })
    expect(search).toBeInTheDocument()
    expect(search.classList.toString()).toEqual('col-start-5 text-lg content-center text-center hover:bg-gray-800')
    expect(search.getElementsByTagName('span')[0].classList.toString()).toEqual('border-b-4 border-blue-500 py-6')
  })
})
