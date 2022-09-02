import { useState } from 'react'
import Navbar from "./components/Navbar"
import HomePage from "./components/HomePage/Home"
import Votings from "./components/HomePage/Votings"
import AboutPage from './components/AboutPage/AboutPage'
import CreatePage from './components/CreatePage/CreatePage'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Footer from './components/Footer'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div>
      <Navbar></Navbar>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/votings" element={<Votings />} />
          <Route path="/about" element={<AboutPage />} />
          <Route path="/create" element={<CreatePage />} />
        </Routes>
      </BrowserRouter>
      <Footer />
    </div>
  )
}

export default App
