import UnfollowButton from '@/components/unfollow-button'
import '@testing-library/jest-dom'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'

describe('UnfollowButton', () => {
  it('フォローボタンをマウスホバーするとフォロー解除になり、押すことができる', async () => {
    const mockOnClick = jest.fn()
    render(<UnfollowButton onClick={mockOnClick} />)

    const followingButton = screen.getByRole('button', { name: 'フォロー中' })

    expect(followingButton).toBeInTheDocument()

    await waitFor(async () => {
      await userEvent.hover(followingButton)
    })

    const unfollowButton = screen.getByRole('button', { name: 'フォロー解除' })

    expect(unfollowButton).toBeInTheDocument()
    await waitFor(async () => {
      await userEvent.click(unfollowButton)
    })

    expect(mockOnClick).toHaveBeenCalledTimes(1)
  })
})
