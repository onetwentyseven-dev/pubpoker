<template>
  <Layout>
    <b-container class="top mb-4">
      <b-row>
        <b-col>
          <div v-if="loading" class="d-flex justify-content-center">
            <b-spinner style="width: 2rem; height: 2rem" />
            <h3 class="ms-3">Loading Tournaments...</h3>
          </div>
          <div v-else>
            <h1>PPC Tournament Registration</h1>
            <hr class="mt-0" />
          </div>
        </b-col>
      </b-row>
      <b-row v-if="!loading">
        <b-col lg="6">
          <b-card header="Create New Tournament" header-tag="header">
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
              <b-button
                variant="primary"
                v-show="selectedVenue"
                @click="createTournment"
                >Create Tournament</b-button
              >
            </div>
          </b-card>
        </b-col>
        <b-col>
          <b-card
            header="Existing Tournaments (Click to select)"
            header-tag="header"
            no-body
          >
            <b-list-group flush>
              <b-list-group-item
                v-for="tournament in tournaments"
                :key="tournament.id"
                :href="`/tournaments/${tournament.id}`"
              >
                {{ fmtTime(tournament.createdAtDate) }}
                {{ fmtVenue(tournament.venueID) }}
              </b-list-group-item>
            </b-list-group>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
  </Layout>
</template>
<script>
import Layout from "../layout/index.vue";
import Error from "../components/Error";
import { DateTime } from "luxon";

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
      tournaments: [],
      error: "",
    };
  },
  methods: {
    fmtTime(t) {
      return DateTime.fromISO(t).toLocaleString(DateTime.DATE_MED_WITH_WEEKDAY);
    },
    fmtVenue(vID) {
      const { name = "Unknown Venue" } =
        this.venues.find((v) => v.id === vID) || {};
      return name;
    },
    async createTournment() {
      await this.axios
        .post(`${apiURL}/tournaments`, {
          venueID: this.selectedVenue.id,
        })
        .then((res) => {
          this.$router.push({
            name: "tournaments",
            params: { tournamentID: res.data.id },
          });
        });
    },
    getVenues() {
      return this.axios
        .get(`${apiURL}/venues`)
        .then((res) => (this.venues = res.data))
        .catch((err) => {
          console.log("failed to load venues", err.response.data.error);
          this.error = "Unable to load venues, please try again later";
        });
    },
    getTournaments() {
      return this.axios
        .get(`${apiURL}/tournaments`)
        .then((res) => (this.tournaments = res.data))
        .catch((err) => {
          console.log("failed to load tournaments", err.response.data.error);
          this.error = "Unable to load tournaments, please try again later";
        });
    },
  },
  async created() {
    this.loading = true;
    await Promise.all([this.getVenues(), this.getTournaments()]).then(() => {
      this.loading = false;
    });
  },
};
</script>
<style>
.top {
  margin-top: 100px;
}
</style>