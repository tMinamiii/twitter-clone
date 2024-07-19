type Props = {
  setSearchQuery: Function
}

export default function SearchInput({ setSearchQuery }: Props) {

  return (
    <input
      type='text'
      className='p-2.5 rounded-full text-black text-lg w-full outline-none text-gray-100 bg-gray-800'
      onChange={async (e) => {
        setSearchQuery(e.target.value)
      }}
    />
  )
}
