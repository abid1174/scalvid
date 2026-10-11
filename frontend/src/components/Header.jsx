export default function Header({ search, onSearchChange, onAddClick }) {
  return (
    <header className="header">
      <div className="logo">Scalvid</div>

      <input
        className="search"
        type="search"
        placeholder="Search videos"
        aria-label="Search videos"
        value={search}
        onChange={(e) => onSearchChange(e.target.value)}
      />

      <button className="btn btn-primary" onClick={onAddClick}>
        + Add video
      </button>
    </header>
  )
}
