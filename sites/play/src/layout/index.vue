<template>
  <div>
    <Navbar />
    <div style="margin-bottom: 15vh">
      <slot></slot>
    </div>
    <div class="footer">
      <b-container>
        <b-row>
          <b-col>
            &copy;{{ thisYear }} Pub Poker Championship
            <br />
            All Rights Reserved.
          </b-col>
          <b-col class="text-right">
            Website designed with
            <b-icon icon="suit-heart-fill" rotate="90" /> by ECH Consulting Inc.
          </b-col>
        </b-row>
      </b-container>
    </div>
  </div>
</template>

<script>
import Navbar from "../components/Navbar.vue";

export default {
  components: {
    Navbar,
  },
  computed: {
    thisYear() {
      return new Date().getFullYear();
    },
  },
};
</script>

<style scoped>
.footer {
  min-height: 10vh;
  position: fixed;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #999;
  /* color: white; */
  left: 0;
  bottom: 0;
  width: 100%;
}
</style>
