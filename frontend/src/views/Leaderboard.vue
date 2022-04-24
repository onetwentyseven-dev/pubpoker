<template>
  <Layout>
    <b-container class="marginTop">
      <div v-if="loading">
        <div class="d-flex justify-content-center mb-3">
          <b-spinner
            style="width: 3rem; height: 3rem"
            label="Loading Leaderboards..."
          />
          <h1 class="ms-3">Loading Leaderboards...</h1>
        </div>
      </div>
      <div v-else-if="!loading">
        <b-row>
          <b-col md="6">
            <h2>Leaderboard {{ leaderboard.length }}</h2>
            <hr />
            <b-card no-body>
              <b-card-body class="searchField" style="background-color: #eee">
                <b-row>
                  <b-col lg="3" class="d-flex">
                    <div class="align-self-center">Name</div>
                  </b-col>
                  <b-col>
                    <b-form-input
                      id="name-search"
                      v-model="searchValue"
                      class="align-self-center"
                      debounce="500"
                      placeholder="Type Player Name to Search Specific Player"
                    />
                  </b-col>
                </b-row>
              </b-card-body>
              <b-table
                :items="leaderboard.items"
                :fields="leaderboard.fields"
                class="m-0"
              >
                <template #cell(points_awarded)="data">
                  {{ fmtPoints(data.value) }}
                </template>
              </b-table>
            </b-card>
          </b-col>
          <b-col md="6">
            <h2>Recent Tournament Winners</h2>
            <hr />
            <b-list-group>
              <b-list-group-item
                v-for="(winner, index) in winners"
                :key="index"
              >
                <b-media>
                  <template #aside>
                    <b-img
                      :style="
                        'background-image: url(' +
                        winner.Player.PlayerSetting.avatar_url +
                        ');'
                      "
                      class="player-avater rounded"
                    />
                  </template>
                  <h5>{{ winner.Player.name }}</h5>
                  <small>{{ winner.Tournament.Venue.name }}</small>
                </b-media>
              </b-list-group-item>
            </b-list-group>
          </b-col>
        </b-row>
      </div>
    </b-container>
  </Layout>
</template>
<script>
import Layout from "../layout/index.vue";

const apiURL = "https://api.ppc.onetwentyseven.dev";

export default {
  components: {
    Layout,
  },
  methods: {
    fmtPoints(p) {
      return new Intl.NumberFormat().format(p);
    },

    async handleSearchValue() {
      await this.fetchLeaderboard();
    },
    fetchLeaderboard() {
      let endpoint = `${apiURL}/seasons/${this.season.id}/leaderboard`;
      const params = [];
      if (this.searchValue !== "") {
        params.push(`search=${this.searchValue}`);
      }

      if (params.length > 0) {
        endpoint = `${endpoint}?${params.join("&")}`;
      }

      return this.axios
        .get(endpoint)
        .then((res) => {
          this.leaderboard.items = res.data;
        })
        .catch((err) => {
          console.log(err.response);
        });
    },
    fetchRecentWinners() {
      const endpoint = `${apiURL}/recent-winners`;

      return this.axios
        .get(endpoint)
        .then((res) => {
          this.winners = res.data;
        })
        .catch((err) => {
          console.log(err.response);
        });
    },
  },
  data() {
    return {
      leaderboard: {
        fields: ["ranking", "name", { key: "points_awarded" }],
        items: [],
      },
      winners: [],
      season: {},
      searchValue: "",
      apiURL: apiURL,
      loading: true,
    };
  },
  watch: {
    searchValue: "handleSearchValue",
  },
  async created() {
    this.season = await this.axios
      .get(`${apiURL}/seasons/current`)
      .then((res) => res.data);
    await Promise.all([
      this.fetchLeaderboard(),
      this.fetchRecentWinners(),
    ]).then(() => (this.loading = false));
  },
};
</script>
<style scoped>
.marginTop {
  margin-top: 150px;
}

.header {
  color: white;
  text-shadow: #000 1px 1px 10px;
}

.player-avater {
  width: 80px;
  height: 80px;
  background-size: cover;
  background-position: center center;
  background-repeat: no-repeat;
}
</style>
