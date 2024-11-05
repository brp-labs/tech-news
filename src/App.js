import React, { useEffect, useState } from 'react';

const App = () => {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch('http://localhost:8080/api/rss')
      .then(response => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json();
      })
      .then(data => {
        setArticles(data);
        setLoading(false);
      })
      .catch(error => {
        setError(error);
        setLoading(false);
      });
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <div>
      <h1>The Latest News on AI and Machine Learning from TechXplore.com</h1>
      <ul>
        {articles.map((article, index) => (
          <li key={index}>
            <h2>{article.title}</h2>
            <p>{article.description}</p>
            <a href={article.link} target="_blank" rel="noopener noreferrer">Read more</a>
            <p><strong>Category:</strong> {article.category}</p>
            <p><strong>Published:</strong> {article.pubDate}</p>
            {article.thumbnail && <img src={article.thumbnail} alt="Thumbnail" />}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default App;
