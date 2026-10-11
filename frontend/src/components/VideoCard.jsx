const viewsFormatter = new Intl.NumberFormat('en', { notation: 'compact' })

export default function VideoCard({ video }) {
  // Stable per-video hue so placeholder thumbnails are distinguishable
  const hue = (video.id * 47) % 360

  return (
    <article className="card">
      <div
        className="thumbnail"
        style={{ background: `linear-gradient(135deg, hsl(${hue} 60% 45%), hsl(${(hue + 40) % 360} 60% 30%))` }}
      >
        <span className="duration">{video.duration}</span>
      </div>
      <div className="card-body">
        <h3 className="card-title">{video.title}</h3>
        <p className="card-meta">{video.channel}</p>
        <p className="card-meta">
          {viewsFormatter.format(video.views)} views · {video.uploadedAt}
        </p>
      </div>
    </article>
  )
}
