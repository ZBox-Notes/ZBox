import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'
import './index.css'
import Layout from './layout.tsx'
import BoxesPage from './pages/boxes/boxes.tsx'
import Home from './pages/home/home.tsx'
import Note from './pages/notes/note.tsx'
import NotesPage from './pages/notes/notes.tsx'

const queryClient = new QueryClient()

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />} >
            <Route index element={<Home />} />
            <Route path="/notes" element={<NotesPage />} />
            <Route path='/notes/:note_id' element={<Note />} />
            <Route path="/boxes" element={<BoxesPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  </StrictMode>,
)
