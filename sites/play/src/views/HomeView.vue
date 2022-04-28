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
            <div v-if="loading">
              <b-spinner style="width: 2rem; height: 2rem" />
              <h4 class="ms-3">Loading Venues...</h4>
            </div>
            <div v-else>
              <div v-if="error != ''">
                <Error :text="error" />
              </div>
              <div v-else>
                <b-form-group
                  id="tournament-input-group"
                  label="Search A Venue:"
                  label-for="search-input"
                >
                  <vue-bootstrap-typeahead
                    :data="venues"
                    v-model="venueSearch"
                    size="lg"
                    :serializer="(s) => s.name"
                    @hit="selectedVenue = $event"
                    :minMatchingChars="1"
                    id="search-input"
                  />
                </b-form-group>
                <b-button v-show="selectedVenue" @click="createTournment"
                  >Create Tournament</b-button
                >
              </div>
            </div>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </Layout>
</template>
<script>
import Layout from "../layout/index.vue";
import Error from "../components/Error";

const apiURL = "https://api.ppc.onetwentyseven.dev";

export default {
  components: {
    Layout,
    Error,
  },
  data() {
    return {
      loading: true,
      venueSearch: "",
      selectedVenue: null,
      venues: [],
      error: "",
    };
  },
  methods: {
    async createTournment() {
      const date = new Date();
      console.log(date);
      // await this.axios.post(`${apiURL}/tournaments`, {

      // }).then((res) => {
      //   console.log(res.data);
      // });
    },
  },
  async created() {
    await this.axios
      .get(`${apiURL}/venues`)
      .then((res) => (this.venues = res.data))
      .catch((err) => {
        console.log("failed to load venues", err.response.data.error);
        this.error = "Unable to load venues, please try again later";
      });
    this.loading = false;
  },
};
</script>
<style>
.top {
  margin-top: 100px;
}
</style>