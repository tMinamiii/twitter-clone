import FollowButton from '@/components/follow-button'
import '@testing-library/jest-dom'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'

describe('FollowButton', () => {
  it('フォローボタンが表示されていて、押すことができる', async () => {
    const mockOnClick = jest.fn()
    render(<FollowButton onClick={mockOnClick} />)

    const button = screen.getByRole('button', { name: 'フォロー' })
    expect(button).toBeInTheDocument()
    await waitFor(async () => {
      await userEvent.click(button)
    })

    expect(mockOnClick).toHaveBeenCalledTimes(1)
  })
})
