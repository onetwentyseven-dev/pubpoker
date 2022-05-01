<template>
  <b-navbar toggleable="lg" type="dark" variant="dark" fixed="top">
    <b-container>
      <b-navbar-brand to="/">
        <img class="logo" src="../assets/spade.svg" />
      </b-navbar-brand>
      <!-- <b-navbar-toggle target="nav-collapse"></b-navbar-toggle> -->

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item href="/">Tournaments</b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-container>
  </b-navbar>
</template>

<script>
import { mapGetters, mapActions } from "vuex";

export default {
  data() {
    return {
      atTop: true,
    };
  },
  methods: {
    ...mapGetters(["getTop"]),
    ...mapActions(["setAtTop"]),
    handleScrollEvent(event) {
      const scrollTop = event.target.scrollingElement.scrollTop;
      const at = scrollTop <= 100;
      this.atTop = at;
      this.setAtTop(at);
    },
  },

  beforeDestroy() {
    window.removeEventListener("scroll", this.handleScrollEvent);
  },
  created() {
    window.addEventListener("scroll", this.handleScrollEvent);
  },
};
</script>

<style scoped>
.logo {
  height: 50px;
  transition: all 0.5s;
}
</style>
