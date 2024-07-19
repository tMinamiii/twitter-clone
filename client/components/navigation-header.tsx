import { Page } from '@/const/navigation'
import Image from 'next/image'
import Link from 'next/link'
import iconImage from '../public/tweet.png'

type Props = {
  page: Page
}

export default function NavigationHeader({ page }: Props) {
  const selected = (p: Page): string => {
    return p === page ? 'border-b-4 border-blue-500 py-6' : 'border-b-0'
  }
  return (
    <div className='grid grid-cols-subgrid grid-cols-5'>
      <div className='flex justify-center items-center m-4'>
        <Image src={iconImage} height={50} width={50} alt='' />
      </div>
      <Link href='/' className='col-start-4 text-lg content-center text-center hover:bg-gray-800'>
        <span className={`${selected(Page.Timeline)}`}>Timeline</span>
      </Link >
      <Link href='/search' className='col-start-5 text-lg content-center text-center hover:bg-gray-800'>
        <span className={`${selected(Page.Search)}`}>Search</span>
      </Link>
    </div>
  )
}
