import Timeline from '@/app/page'
import '@testing-library/jest-dom'
import { act, render, screen } from '@testing-library/react'

jest.mock('../../components/actions.ts', () => ({
  getDummySession: async () => { },
}))

describe('Timeline', () => {
  it('タイムライン全体が表示される', async () => {
    global.fetch = jest.fn().mockImplementation(() => {
      return new Promise((resolve) => {
        resolve({
          ok: true,
          json: () => {
            return {
              count: 2,
              lastUuid: 2,
              posts: [
                {
                  uuid: '1',
                  accountId: 'test_account_id_01',
                  username: 'test_name_01',
                  content: 'tweet tweet tweet',
                  postedAt: '2024-06-13T11:22:33Z',
                },
                {
                  uuid: '2',
                  accountId: 'test_account_id_02',
                  username: 'test_name_02',
                  content: 'hello',
                  postedAt: '2024-06-14T11:22:33Z',
                },
              ]
            }
          }
        })
      })
    })


    await act(async () => { render(<Timeline />) })

    expect(screen.getByText('Timeline')).toBeInTheDocument()
    expect(screen.getByText('Search')).toBeInTheDocument()

    expect(screen.getByRole('textbox')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: 'Post' })).toBeInTheDocument

    expect(screen.getByText('test_name_01')).toBeInTheDocument()
    expect(screen.getByText('@test_account_id_01')).toBeInTheDocument()
    expect(screen.getByText('2024年06月13日')).toBeInTheDocument()
    expect(screen.getByText('tweet tweet tweet')).toBeInTheDocument()

    expect(screen.getByText('test_name_02')).toBeInTheDocument()
    expect(screen.getByText('@test_account_id_02')).toBeInTheDocument()
    expect(screen.getByText('2024年06月14日')).toBeInTheDocument()
    expect(screen.getByText('hello')).toBeInTheDocument()
  })
})
