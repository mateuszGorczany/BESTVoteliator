import { useState } from 'react'
import Navbar from "./components/Navbar"
import HomePage from "./components/HomePage/Home"
import Voting from "./components/VotingPage/Voting"
import AboutPage from './components/AboutPage/AboutPage'
import CreatePage from './components/CreatePage/CreatePage'
import { HashRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import Footer from './components/Footer'
import { useEffect } from 'react'
import { initLogin } from "./utils"

function App() {
  const [count, setCount] = useState(0)
  useEffect(() => initLogin(), [])
  return (
    <div>
      <Navbar />
      <HashRouter>
        <Routes>
          <Route exact path="/" element={<HomePage />} />
          <Route exact path="/voting" element={<Voting />} />
          <Route exact path="/about" element={<AboutPage />} />
          <Route exact path="/create" element={<CreatePage />} />
        </Routes>
      </HashRouter>
      <Footer />
    </div>
  )
}

export default App
