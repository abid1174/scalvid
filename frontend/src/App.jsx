import { useMemo, useState } from 'react'
import Header from './components/Header.jsx'
import VideoCard from './components/VideoCard.jsx'
import AddVideoModal from './components/AddVideoModal.jsx'
import { initialVideos } from './data/videos.js'

export default function App() {
  const [videos, setVideos] = useState(initialVideos)
  const [search, setSearch] = useState('')
  const [isAddOpen, setIsAddOpen] = useState(false)

  const filteredVideos = useMemo(() => {
    const query = search.trim().toLowerCase()
    if (!query) return videos
    return videos.filter(
      (v) => v.title.toLowerCase().includes(query) || v.channel.toLowerCase().includes(query),
    )
  }, [videos, search])

  function handleAdd({ title, channel }) {
    const newVideo = {
      id: Date.now(),
      title,
      channel,
      views: 0,
      duration: '00:00',
      uploadedAt: new Date().toISOString().slice(0, 10),
    }
    setVideos((prev) => [newVideo, ...prev])
    setIsAddOpen(false)
  }

  return (
    <>
      <Header search={search} onSearchChange={setSearch} onAddClick={() => setIsAddOpen(true)} />

      <main className="content">
        {filteredVideos.length > 0 ? (
          <div className="grid">
            {filteredVideos.map((video) => (
              <VideoCard key={video.id} video={video} />
            ))}
          </div>
        ) : (
          <p className="empty">No videos match “{search}”.</p>
        )}
      </main>

      <AddVideoModal open={isAddOpen} onClose={() => setIsAddOpen(false)} onAdd={handleAdd} />
    </>
  )
}
