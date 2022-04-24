<template>
  <div>
    <Navbar />
    <slot></slot>
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
  created() {
    console.log("Layout");
  },
};
</script>

<style scoped>
.footer {
  min-height: 10vh;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #666;
  color: white;
}
</style>
