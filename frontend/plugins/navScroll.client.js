export default ({ store }) => {
  window.addEventListener('scroll', (event) => {
    return handleScrollEvent(event, store)
  })

  return () => {
    window.removeEventListener('scroll', (event) => {
      return handleScrollEvent(event, store)
    })
  }
}

function handleScrollEvent(event, store) {
  const scrollTop = event.target.scrollingElement.scrollTop
  store.dispatch('setAtTop', scrollTop <= 100)
}
