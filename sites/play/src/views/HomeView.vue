<template>
  <Layout>
    <b-container class="top mb-4">
      <b-row>
        <b-col>
          <h1>PPC Tournament Registration</h1>
          <hr class="mt-0" />
        </b-col>
      </b-row>
      <b-row>
        <b-col lg="6">
          <b-card header="Create New Tournament" header-tag="header">
            <b-card-text>Header and footers using props.</b-card-text>
            <b-button href="#" variant="primary">Go somewhere</b-button>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </Layout>
</template>
<script>
import Layout from "../layout/index.vue";

export default {
  components: {
    Layout,
  },
  data() {
    return {};
  },
};
</script>
<style>
.top {
  margin-top: 100px;
}
</style>