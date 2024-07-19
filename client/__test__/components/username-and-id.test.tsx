import UsernameAndID from '@/components/username-and-id'
import '@testing-library/jest-dom'
import { render, screen } from '@testing-library/react'

describe('UsernameAndID', () => {
  it('propしたユーザー名と@アカウントIDが表示されている', async () => {
    render(<UsernameAndID username='test_name_01' accountId='test_account_id_01' />)

    const username = screen.getByText('test_name_01')

    expect(username).toBeInTheDocument()

    const accountId = screen.getByText('@test_account_id_01')

    expect(accountId).toBeInTheDocument()
  })
})
