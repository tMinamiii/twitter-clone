import { useState } from 'react'

type Props = {
  onClick: () => Promise<void>
}

export default function UnfollowButton({ onClick }: Props) {
  const [isHovering, setIsHovered] = useState(false)
  const onMouseEnter = () => setIsHovered(true)
  const onMouseLeave = () => setIsHovered(false)

  return (
    <div onMouseEnter={onMouseEnter} onMouseLeave={onMouseLeave} >
      {!isHovering
        ? <button
          className='bg-black-100 text-white text-sm py-2 px-4 rounded-full border-solid border border-white'
          type='button'
        >
          フォロー中
        </button>
        : <button
          className='bg-black-100 text-red-500 text-sm py-2 px-4 rounded-full border-solid border border-red-500'
          type='button'
          onClick={onClick}>
          フォロー解除
        </button>
      }
    </div>
  )
}
