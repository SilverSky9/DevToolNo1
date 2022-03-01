import '../styles/Home.module.css'
import { Row } from 'react-bootstrap';

const Layout = ({ children }) => {

  return (<div className="container-fluid " style={{ backgroundColor: "#fefefe" }} >

    <Row>
      <div className="col-sm-auto p-0 position-fixed h-100 bg-light font">
        <ul
          className="
              nav nav-pills nav-flush
              flex-sm-column flex-row flex-nowrap
              mb-auto
              mx-4
              mt-2
              text-center
              "
          style={{ fontFamily: 'Raleway' }}
        >
          <li className="pl-2  ">
            <div className='display-3' > Mar<span style={{ color: '#197DFF' }}>ket</span></div>

          </li>
          <li
            className="pl-2 my-4"
          >
            search

          </li>
        </ul>
      </div>
      <div className="col-2 "  >
      </div>
      <div className="col-10 p-0 " style={{ fontFamily: 'Raleway' }}  >
        {children}

      </div>
    </Row>
  </div>);
}

export default Layout;