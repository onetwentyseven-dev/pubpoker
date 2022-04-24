<template>
  <Layout>
    <b-container class="marginTop">
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
      </b-row>
    </b-container>
  </Layout>
</template>
<script>
import Layout from "../layout/index";
import Error from "../components/Error";

const apiURL = "https://api.ppc.onetwentyseven.dev";

export default {
  components: {
    Layout,
    Error,
  },
  methods: {
    fmtPoints(p) {
      return new Intl.NumberFormat().format(p);
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
          this.error = "Failed to load leaderboard. Please try again later.";
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
          this.error = "Failed to load Recent Winners. Please try again later";
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
      error: "",
    };
  },
  watch: {
    searchValue: "fetchLeaderboard",
  },
  async created() {
    try {
      await this.axios.get(`${apiURL}/seasons/current`).then((res) => {
        this.season = res.data;
      });
      await Promise.all([this.fetchLeaderboard(), this.fetchRecentWinners()]);
    } catch (err) {
      console.log(err.response.data.error);
      this.error =
        "One or more requets needed to initialize the leaderboards failed to return. Please try again later";
    }
    this.loading = false;
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
