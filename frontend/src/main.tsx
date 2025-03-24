import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Route, Routes } from 'react-router'
import './index.css'
import Layout from './layout.tsx'
import Boxes from './pages/boxes/boxes.tsx'
import Home from './pages/home/home.tsx'
import Notes from './pages/notes/notes.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />} >
          <Route index element={<Home />} />
          <Route path="/notes" element={<Notes />} />
          <Route path="/boxes" element={<Boxes />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
