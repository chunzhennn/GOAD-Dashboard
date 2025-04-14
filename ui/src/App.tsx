import React, { useState, useMemo } from 'react';
import { Provider } from 'react-redux';
import { configureStore } from '@reduxjs/toolkit';
import { api } from './store/api';
import { Layout } from './components/layout/Layout';
import { Dashboard } from './pages/Dashboard';
import { ThemeContext, ThemeType } from './context/ThemeContext';

// Theme provider component
const ThemeProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  // Initialize from localStorage or fallback to 'system'
  const [theme, setTheme] = useState<ThemeType>(() => {
    const savedTheme = localStorage.getItem('theme');
    // Using nullish coalescing to handle null case
    return (savedTheme as ThemeType) ?? 'system';
  });

  // Apply theme to document
  React.useEffect(() => {
    const applyTheme = () => {
      if (theme === 'dark' || (theme === 'system' && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
    };

    applyTheme();
    localStorage.setItem('theme', theme);

    // Add listener for system theme changes
    if (theme === 'system') {
      const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
      const handleChange = () => { applyTheme(); };
      mediaQuery.addEventListener('change', handleChange);
      return () => { mediaQuery.removeEventListener('change', handleChange); };
    }
  }, [theme]);
  
  // Memoize the context value to prevent unnecessary re-renders
  const contextValue = useMemo(() => ({ theme, setTheme }), [theme]);

  return (
    <ThemeContext.Provider value={contextValue}>
      {children}
    </ThemeContext.Provider>
  );
};

// Configure Redux store
const store = configureStore({
  reducer: {
    [api.reducerPath]: api.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(api.middleware),
});

function App() {
  return (
    <ThemeProvider>
      <Provider store={store}>
        <Layout>
          <Dashboard />
        </Layout>
      </Provider>
    </ThemeProvider>
  );
}

export default App;
