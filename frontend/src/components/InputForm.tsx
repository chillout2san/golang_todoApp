export const InputForm = () => {
  return (
    <div className="w-auto h-30 mb-4 p-4 border border-gray-200 rounded shadow-lg">
      <p className="mb-2 font-bold">新しいタスクを追加する</p>
      <input
        placeholder="買い物"
        className="mr-4 border shadow-md border-teal-500 rounded"
      />
      <button className="px-2 h-6 border border-white rounded bg-teal-400 shadow-md text-white">
        追加
      </button>
    </div>
  )
}
