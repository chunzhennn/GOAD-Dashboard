import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

// Initialize dark mode based on localStorage or system preference
const initializeDarkMode = () => {
  const savedTheme = localStorage.getItem('theme');
  
  if (savedTheme === 'dark') {
    // Use dark mode if explicitly saved
    document.documentElement.classList.add('dark');
  } else if (savedTheme === 'light') {
    // Use light mode if explicitly saved
    document.documentElement.classList.remove('dark');
  } else {
    // Default to system preference if no saved preference or set to 'system'
    const isDarkMode = window.matchMedia('(prefers-color-scheme: dark)').matches;
    if (isDarkMode) {
      document.documentElement.classList.add('dark');
    }
  }
};

// Call the function before rendering
initializeDarkMode();

const rootElement = document.getElementById('root');
if (rootElement) {
  ReactDOM.createRoot(rootElement).render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  );
} else {
  console.error('Root element not found in the document');
}
