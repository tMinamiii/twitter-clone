import SearchInput from '@/app/search/search-input'
import '@testing-library/jest-dom'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'

describe('SearchInput', () => {
  it('検索語入力inputが表示され、テキスト入力ができる', async () => {
    const mockSetSearchQuery = jest.fn()
    render(
      <SearchInput setSearchQuery={mockSetSearchQuery} />
    )

    const input = screen.getByRole('textbox')
    expect(input).toBeInTheDocument()

    await waitFor(async () => {
      await userEvent.type(input, 'abc')
    })

    expect(screen.getByDisplayValue('abc')).toBeInTheDocument()

    expect(mockSetSearchQuery).toHaveBeenCalledTimes(3)
  })
})
