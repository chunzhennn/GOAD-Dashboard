import { useTheme } from '../hooks/useTheme';

export function ThemeToggle() {
  const { theme, setTheme } = useTheme();

  return (
    <div className="flex items-center rounded-lg bg-gray-100 dark:bg-gray-800 p-1">
      {/* Light theme option */}
      <button
        type="button"
        onClick={() => {
          setTheme('light');
        }}
        className={`flex items-center justify-center px-3 py-1.5 text-sm font-medium rounded-md transition-colors ${
          theme === 'light' 
            ? 'bg-white dark:bg-gray-700 shadow-sm' 
            : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200'
        }`}
        aria-label="Light theme"
      >
        <svg 
          xmlns="http://www.w3.org/2000/svg" 
          width="16" 
          height="16" 
          viewBox="0 0 24 24" 
          fill="none" 
          stroke="currentColor" 
          strokeWidth="2" 
          strokeLinecap="round" 
          strokeLinejoin="round"
          className="mr-1"
        >
          <circle cx="12" cy="12" r="5" />
          <path d="M12 1v2M12 21v2M4.2 4.2l1.4 1.4M18.4 18.4l1.4 1.4M1 12h2M21 12h2M4.2 19.8l1.4-1.4M18.4 5.6l1.4-1.4" />
        </svg>
        Light
      </button>

      {/* Dark theme option */}
      <button
        type="button"
        onClick={() => {
          setTheme('dark');
        }}
        className={`flex items-center justify-center px-3 py-1.5 text-sm font-medium rounded-md transition-colors ${
          theme === 'dark' 
            ? 'bg-white dark:bg-gray-700 shadow-sm' 
            : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200'
        }`}
        aria-label="Dark theme"
      >
        <svg 
          xmlns="http://www.w3.org/2000/svg" 
          width="16" 
          height="16" 
          viewBox="0 0 24 24" 
          fill="none" 
          stroke="currentColor" 
          strokeWidth="2" 
          strokeLinecap="round" 
          strokeLinejoin="round"
          className="mr-1"
        >
          <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z" />
        </svg>
        Dark
      </button>

      {/* System theme option */}
      <button
        type="button"
        onClick={() => {
          setTheme('system');
        }}
        className={`flex items-center justify-center px-3 py-1.5 text-sm font-medium rounded-md transition-colors ${
          theme === 'system' 
            ? 'bg-white dark:bg-gray-700 shadow-sm' 
            : 'text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200'
        }`}
        aria-label="System theme"
      >
        <svg 
          xmlns="http://www.w3.org/2000/svg" 
          width="16" 
          height="16" 
          viewBox="0 0 24 24" 
          fill="none" 
          stroke="currentColor" 
          strokeWidth="2" 
          strokeLinecap="round" 
          strokeLinejoin="round"
          className="mr-1"
        >
          <rect width="20" height="14" x="2" y="3" rx="2" />
          <line x1="8" x2="16" y1="21" y2="21" />
          <line x1="12" x2="12" y1="17" y2="21" />
        </svg>
        System
      </button>
    </div>
  );
} 