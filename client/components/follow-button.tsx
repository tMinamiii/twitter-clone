type Props = {
  onClick: () => Promise<void>
}

export default function FollowButton({ onClick }: Props) {
  return (
    <button
      className="bg-gray-100 hover:bg-gray-400 text-black text-sm py-2 px-4 border rounded-full"
      type="submit"
      onClick={onClick}>
      フォロー
    </button>
  )
}
