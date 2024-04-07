import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [books, setBooks] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/books')  // Adjust this URL to where your backend is running
      .then(response => response.json())
      .then(data => setBooks(data))
      .catch(err => console.error('Error fetching books:', err));
  }, []); // Empty dependency array means this effect runs once after the initial render

  return (
    <div className="App">
      <header className="App-header">
        <h1>Books List</h1>
        <ul>
          {books.map(book => (
            <li key={book.id}>
              {book.title} by {book.author} (Published on: {new Date(book.published_date).toLocaleDateString()})
            </li>
          ))}
        </ul>
      </header>
    </div>
  );
}

export default App;
