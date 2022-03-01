import '../styles/globals.css'
import 'bootstrap/dist/css/bootstrap.css'
import '../styles/Home.module.css'
import { useRouter } from 'next/router'

import Layout from '../components/Layout'

function MyApp({ Component, pageProps }) {
  const router = useRouter()
  if (router.asPath === '/') {
    return (
      <Component {...pageProps} />
    )
  }
  return (
    <Layout>
      <Component {...pageProps} />

    </Layout>
  )
}

export default MyApp
