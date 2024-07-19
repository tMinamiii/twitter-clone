import Tweets from '@/app/tweets'
import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react'

describe('Tweets', () => {
  it('Propsの内容通りつぶやき一覧が表示される', () => {
    render(<Tweets tweets={
      [
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
    } />)

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
