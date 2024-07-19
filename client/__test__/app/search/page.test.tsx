import Search from '@/app/search/page'
import '@testing-library/jest-dom'
import { act, render, screen } from '@testing-library/react'

jest.mock('../../../components/actions.ts', () => ({
  getDummySession: async () => { },
}))

describe('Page', () => {
  it('検索ページ全体が描画される', async () => {
    global.fetch = jest.fn().mockImplementation(() => {
      return new Promise((resolve) => {
        resolve({
          ok: true,
          json: () => {
            return {
              count: 2,
              lastUuid: '1',
              users: [
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
              ],
            }
          }
        })
      })
    })


    await act(async () => { render(<Search />) })

    expect(screen.getByText('Timeline')).toBeInTheDocument()
    expect(screen.getByText('Search')).toBeInTheDocument()

    expect(screen.getByRole('textbox')).toBeInTheDocument()

    expect(screen.getByText('test_name_01')).toBeInTheDocument()
    expect(screen.getByText('@test_account_id_01')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: 'フォロー' })).toBeInTheDocument()

    expect(screen.getByText('test_name_02')).toBeInTheDocument()
    expect(screen.getByText('@test_account_id_02')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: 'フォロー中' }))
  })
})
