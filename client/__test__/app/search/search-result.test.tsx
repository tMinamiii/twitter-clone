import SearchResult from '@/app/search/search-result'
import '@testing-library/jest-dom'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'

jest.mock('../../../components/actions.ts', () => ({
  getDummySession: async () => { },
}))

describe('SearchResult', () => {
  it('検索結果のユーザー一覧のフォローボタンを押す', async () => {
    global.fetch = jest.fn().mockImplementation(() => {
      return new Promise((resolve) => {
        resolve({
          ok: true,
          json: () => { return {} }
        })
      })
    })

    const mockUpdateResult = jest.fn()

    render(
      <SearchResult updateResult={mockUpdateResult} searchResult={
        [
          {
            username: 'test_name_01',
            accountId: 'test_account_id_01',
            isFollowed: false
          },
          {
            username: 'test_name_02',
            accountId: 'test_account_id_02',
            isFollowed: true
          }
        ]
      } />
    )

    expect(screen.getByText('test_name_01')).toBeInTheDocument()
    expect(screen.getByText('@test_account_id_01')).toBeInTheDocument()

    expect(screen.getByText('test_name_02')).toBeInTheDocument()
    expect(screen.getByText('@test_account_id_02')).toBeInTheDocument()

    const followButton = screen.getByRole('button', { name: 'フォロー' })
    expect(followButton).toBeInTheDocument()
    await waitFor(async () => {
      // call updateResult
      await userEvent.click(followButton)
    })

    expect(mockUpdateResult).toHaveReturnedTimes(1)
  })
  it('検索結果のユーザー一覧のフォロー解除ボタンを押す', async () => {
    global.fetch = jest.fn().mockImplementation(() => {
      return new Promise((resolve) => {
        resolve({
          ok: true,
          json: () => { return {} }
        })
      })
    })

    const mockUpdateResult = jest.fn()
    render(
      <SearchResult updateResult={mockUpdateResult} searchResult={
        [
          {
            username: 'test_name_01',
            accountId: 'test_account_id_01',
            isFollowed: false
          },
          {
            username: 'test_name_02',
            accountId: 'test_account_id_02',
            isFollowed: true
          }
        ]
      } />
    )

    expect(screen.getByText('test_name_01')).toBeInTheDocument()
    expect(screen.getByText('@test_account_id_01')).toBeInTheDocument()

    expect(screen.getByText('test_name_02')).toBeInTheDocument()
    expect(screen.getByText('@test_account_id_02')).toBeInTheDocument()

    const followingButton = screen.getByRole('button', { name: 'フォロー中' })
    expect(followingButton).toBeInTheDocument()

    await waitFor(async () => {
      await userEvent.hover(followingButton)
    })

    const unfollowButton = screen.getByRole('button', { name: 'フォロー解除' })
    expect(unfollowButton).toBeInTheDocument()
    await waitFor(async () => {
      // call updateResult
      await userEvent.click(unfollowButton)
    })

    expect(mockUpdateResult).toHaveReturnedTimes(1)
  })
})
