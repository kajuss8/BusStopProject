import Schedule from './Schedule.jsx'
import {Route, Routes} from "react-router-dom"
import Home from './pages/Home.jsx'
import RouteSchedule from './pages/RouteSchedule.jsx'

function App() {
  return(
    <>
      <div>
        <Routes> 
          <Route path='/' element={<Home/>}/>
          <Route path='/RouteSchedule' element={<RouteSchedule/>}/>
        </Routes>
      </div>
    </>
  )
}

export default App;
