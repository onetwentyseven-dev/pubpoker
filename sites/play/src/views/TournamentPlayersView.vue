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
            <h1>Final Table - {{ this.tournament.venue.name }}</h1>
            <hr class="mt-0" />
            <div v-if="error != ''">
              <Error :text="error" />
            </div>
          </div>
        </b-col>
      </b-row>
      <b-row>
        <b-col lg="6">
          <b-card header="Registered Players" no-body class="mb-2">
            <b-card-body>
              When a player goes out, locate them below and update the number to
              the position the player went out at. The list is in alphabetical
              order, the list on the right will update the final stadings. When
              the tournament is finish, submit the standings for approval.
            </b-card-body>
            <b-table-simple>
              <tr v-for="(player, i) in players" :key="player.id">
                <td>
                  {{ player.name }}
                </td>
                <td>
                  <b-form-input
                    type="number"
                    v-model="players[i].finalPosition"
                    :value="player.finalPosition"
                  />
                </td>
                <td>
                  <b-button @click="updateFinalPosition(player)">Save</b-button>
                </td>
              </tr>
            </b-table-simple>
          </b-card>
        </b-col>
        <b-col lg="6">
          <b-card
            header="Final Standings "
            no-body
            class="mb-2"
            v-show="standings.length > 0"
          >
            <b-table-simple class="m-0">
              <tr v-for="player in standings" :key="player.id">
                <td>
                  {{ player.finalPosition }}
                </td>
                <td>
                  {{ player.name }}
                </td>
              </tr>
            </b-table-simple>
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
      tournament: null,
      venue: null,
      tournamentID: "",
      loading: true,
      players: [],
      standings: [],
      error: "",
    };
  },
  watch: {
    players: function (updated, old) {
      this.getFinalStandings();
    },
  },
  methods: {
    async updateFinalPosition(player) {
      console.log(player);
      await this.axios.patch(
        `${apiURL}/tournaments/${this.tournament.id}/players/${player.playerID}`,
        {
          finalPosition: parseInt(player.finalPosition),
        }
      );
      await this.getTournamentPlayers();
    },
    getFinalStandings() {
      var players =
        this.players.filter((player) => player.finalPosition > 0) || [];
      if (players.length > 0) {
        players = players.sort((a, b) => a.finalPosition - b.finalPosition);
      }
      this.standings = players;
    },
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