import SignIn from './pages/SignIn';
import SignUp from "./pages/SignUp";
import Navbar from "./components/Navbar";
import Hero from './components/Hero';
import Articles from './components/Articles';
import AboutUsHero from "./components/AboutUsHero"
import OurTeam from "./components/OurTeam";
import Footer from "./components/Footer";
import {Routes, Route} from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="/" element={
          <div>
            <Navbar />
            <Hero />
            <Articles />
            <Footer />
          </div>
        } />
        <Route path="/about" element={
          <div>
            <Navbar />
            <AboutUsHero />
            <OurTeam />
            <Footer />
          </div>
        }/>
        <Route path="/signin" element={<SignIn />} />
        <Route path="/signup" element={<SignUp />} />
      </Routes>
    </div>
  );
}

export default App;
