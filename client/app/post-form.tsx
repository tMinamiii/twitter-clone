import PostButton from '@/components/post-button'
import { useRef, useState } from 'react'
import { postTweet } from './actions'

type Props = {
  updateTimeline: () => Promise<void>
}

export default function PostForm({ updateTimeline }: Props) {
  const formRef = useRef<HTMLFormElement>(null)
  const textAreaRef = useRef<HTMLTextAreaElement>(null)
  const [content, setContent] = useState('')

  const onChangeTweetTextArea = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    // adjust textarea height
    if (textAreaRef.current) {
      textAreaRef.current.style.height = 'inherit'
      textAreaRef.current.style.height = `${textAreaRef.current.scrollHeight}px`
    }

    setContent(e.target.value)
  }

  const onKeyDownTweetTextArea = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === 'Enter' && (e.ctrlKey || e.metaKey)) {
      e.preventDefault()
      e.stopPropagation()
      submitPost()
    }
  }

  const submitPost = async () => {
    await postTweet(content)
    await updateTimeline()
    formRef.current?.reset() // textareaをクリア
  }

  return (
    <form ref={formRef}>
      <div>
        <textarea
          className='p-2.5 w-full rounded text-lg bg-black outline-none resize-none'
          placeholder='いまどうしている？'
          rows={1}
          ref={textAreaRef}
          name='content'
          onChange={onChangeTweetTextArea}
          onKeyDown={onKeyDownTweetTextArea}
        />
      </div>
      <div className='text-right p-2'>
        <PostButton onClick={submitPost} />
      </div>
    </form >
  )
}
