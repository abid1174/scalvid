import { useEffect, useRef, useState } from 'react'

export default function AddVideoModal({ open, onClose, onAdd }) {
  const dialogRef = useRef(null)
  const [title, setTitle] = useState('')
  const [channel, setChannel] = useState('')

  useEffect(() => {
    const dialog = dialogRef.current
    if (open && !dialog.open) dialog.showModal()
    if (!open && dialog.open) dialog.close()
  }, [open])

  function handleSubmit(e) {
    e.preventDefault()
    onAdd({ title: title.trim(), channel: channel.trim() })
    setTitle('')
    setChannel('')
  }

  return (
    <dialog ref={dialogRef} className="modal" onClose={onClose}>
      <form onSubmit={handleSubmit}>
        <h2>Add video</h2>

        <label>
          Title
          <input value={title} onChange={(e) => setTitle(e.target.value)} required autoFocus />
        </label>

        <label>
          Channel
          <input value={channel} onChange={(e) => setChannel(e.target.value)} required />
        </label>

        <div className="modal-actions">
          <button type="button" className="btn" onClick={onClose}>
            Cancel
          </button>
          <button type="submit" className="btn btn-primary">
            Add
          </button>
        </div>
      </form>
    </dialog>
  )
}
