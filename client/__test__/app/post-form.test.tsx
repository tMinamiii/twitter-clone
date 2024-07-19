import PostForm from '@/app/post-form'
import '@testing-library/jest-dom'
import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'

jest.mock('../../components/actions.ts', () => ({
  getDummySession: async () => { },
}))

describe('PostForm', () => {
  it('つぶやきを記入して投稿ボタンを押す', async () => {
    global.fetch = jest.fn().mockImplementation(() => {
      return new Promise((resolve) => {
        resolve({
          ok: true,
          json: () => { return {} }
        })
      })
    })

    const mockUpdateTimeline = jest.fn()

    render(<PostForm updateTimeline={mockUpdateTimeline} />)

    const input = screen.getByRole('textbox')
    expect(input).toBeInTheDocument()
    await waitFor(async () => {
      await userEvent.type(input, 'abc')
    })

    expect(screen.getByDisplayValue('abc')).toBeInTheDocument()

    const post = screen.getByRole('button', { name: 'Post' })
    await waitFor(async () => {
      await userEvent.click(post)
    })

    expect(mockUpdateTimeline).toHaveBeenCalledTimes(1)
  })
  it('つぶやきを記入してCtrlEnterで投稿', async () => {
    global.fetch = jest.fn().mockImplementation(() => {
      return new Promise((resolve) => {
        resolve({
          ok: true,
          json: () => { return {} }
        })
      })
    })

    const mockUpdateTimeline = jest.fn()

    render(<PostForm updateTimeline={mockUpdateTimeline} />)

    const input = screen.getByRole('textbox')
    expect(input).toBeInTheDocument()
    await waitFor(async () => {
      await userEvent.type(input, 'abc')
      await userEvent.keyboard('{Control>}{Enter}')
    })
    expect(mockUpdateTimeline).toHaveBeenCalledTimes(1)
  })
})
