<template>
  <Layout>
    <b-container class="top mb-4">
      <b-row>
        <b-col>
          <div v-if="loading" class="d-flex justify-content-center">
            <b-spinner style="width: 2rem; height: 2rem" />
            <h3 class="ms-3">Loading Tournament...</h3>
          </div>
          <div v-else>
            <h1>
              PPC Tournament Registration - {{ this.tournament.venue.name }}
            </h1>
            <hr class="mt-0" />
            <div v-if="error != ''">
              <Error :text="error" />
            </div>
          </div>
        </b-col>
      </b-row>
      <b-row v-if="!loading && error == ''">
        <b-col lg="6">
          <b-card
            header="Player Registration"
            header-tag="header"
            class="mb-2"
            v-show="!this.tournament.isRegistrationClosed"
          >
            <div>
              <b-form-group
                id="tournament-input-group"
                label="Search A Player:"
                label-for="search-input"
              >
                <vue-bootstrap-typeahead
                  :key="inputTracker"
                  :data="results"
                  v-model="playerQuery"
                  :serializer="(s) => s.name"
                  @hit="selectedPlayer = $event"
                  :minMatchingChars="3"
                  @input="searchPlayer"
                  id="search-input"
                  :showAllResults="playerSearchShowAll"
                />
              </b-form-group>
              <b-button
                variant="primary"
                v-show="selectedPlayer"
                @click="registerPlayer"
              >
                Register
              </b-button>
            </div>
          </b-card>
          <b-card header="Tournament Controls" no-body>
            <b-list-group>
              <b-list-group-item @click="handleCloseRegistration" href="#">
                {{ tournament.isRegistrationClosed ? "Open" : "Close" }}
                Registration
              </b-list-group-item>
              <b-list-group-item
                @click="handleInitFinalTable"
                :to="{
                  name: 'players',
                  params: { tournamentID: tournamentID },
                }"
              >
                Start Final Table
              </b-list-group-item>
            </b-list-group>
          </b-card>
        </b-col>
        <b-col>
          <b-card
            :header="`Registered Players (${this.players.length})`"
            header-tag="header"
            no-body
            class="mb-2"
          >
            <b-list-group flush>
              <b-list-group-item v-for="player in players" :key="player.id">
                {{ player.name }}
              </b-list-group-item>
              <b-list-group-item v-if="!players.length">
                No Players have registered
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
import { debounce } from "lodash";

const apiURL = "https://api.ppc.onetwentyseven.dev";

export default {
  components: {
    Layout,
    Error,
  },
  data() {
    return {
      tournament: null,
      tournamentID: "",
      loading: true,
      playerQuery: "hello",
      playerSearchShowAll: false,
      selectedPlayer: null,
      results: [],
      players: [],
      error: "",
      inputTracker: 0,
    };
  },
  methods: {
    async handleCloseRegistration() {
      await this.axios.patch(`${apiURL}/tournaments/${this.tournament.id}`, {
        isRegistrationClosed: !this.tournament.isRegistrationClosed,
      });
      await this.getTournament();
    },
    handleInitFinalTable: () => {},
    async registerPlayer() {
      await this.axios.post(
        `${apiURL}/tournaments/${this.tournament.id}/players`,
        {
          playerID: this.selectedPlayer.id,
        }
      );
      this.inputTracker++;
      this.playerQuery = "";
      await this.getTournamentPlayers();
    },

    searchPlayer: debounce(function () {
      this.axios
        .get(`${apiURL}/players/search?q=${this.playerQuery}`)
        .then((res) => {
          this.results = res.data;
        });
    }, 500),

    getTournament() {
      return this.axios
        .get(`${apiURL}/tournaments/${this.tournamentID}`)
        .then((res) => {
          this.tournament = res.data;
        })
        .catch((err) => {
          console.log("failed to load tournament", err.response.data.error);
          this.error = "Failed to load tournament with provided ID";
        });
    },
    getTournamentVenue() {
      return this.axios
        .get(`${apiURL}/tournaments/${this.tournamentID}/venue`)
        .then((res) => {
          this.venue = res.data;
        });
    },
    getTournamentPlayers() {
      return this.axios
        .get(`${apiURL}/tournaments/${this.tournamentID}/players`)
        .then((res) => {
          this.players = res.data;
        });
    },
  },
  async created() {
    this.tournamentID = this.$router.currentRoute.params.tournamentID;
    await Promise.all([
      this.getTournament(),
      this.getTournamentVenue(),
      this.getTournamentPlayers(),
    ]).then(() => {
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