<template>
  <div>
    <Navbar />
    <Nuxt />
    <div class="footer mt-4">
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
export default {
  computed: {
    thisYear() {
      return new Date().getFullYear()
    },
  },
}
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
