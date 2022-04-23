<template>
  <b-navbar toggleable="lg" type="dark" variant="dark" fixed="top">
    <b-container>
      <b-navbar-brand to="/">
        <b-img
          class="logo"
          :style="atTop ? 'height: 100px' : 'height: 50px'"
          :src="
            atTop ? require('~/assets/logo.svg') : require('~/assets/spade.svg')
          "
        />
      </b-navbar-brand>
      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="ml-auto">
          <b-nav-item href="/">Home</b-nav-item>
          <b-nav-item href="/about">About</b-nav-item>
          <b-nav-item href="/leaderboards">Leaderboards</b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-container>
  </b-navbar>
</template>

<script>
export default {
  computed: {
    atTop() {
      return this.$store.state.atTop
    },
  },
}
</script>

<style scoped>
.logo {
  transition: all 0.5s;
}
</style>
