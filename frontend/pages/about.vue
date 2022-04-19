<template>
  <b-container :class="marginTop">
    <b-row>
      <b-col>
        <h1>About the PPC</h1>
        <hr />
        <p>
          The Pub Poker Championship is a completely FREE poker league with
          tournaments held nightly at the greatest bars, restaurants, and
          taverns in the Charlotte, NC area. Our events are full-service,
          meaning we manage every aspect of our poker tournaments, from our
          fully trained tournament directors, to top quality equipment, to our
          coveted prizes and awards.
        </p>
        <p>
          Over the past 18 years we have become the best and fastest growing
          poker league in our market. Our mission is to provide our players &
          host establishments with an overall experience that exceeds their
          greatest expectation. Our results speak for themselves, but don’t just
          take our word for it… come out and see it for yourself!
        </p>

        <h2>Our League</h2>
        <hr />
        <p>
          We host multiple games almost every night of the week. You register to
          play a game by simply arriving at a tournament location 20 minutes
          prior to the start of the event.
        </p>
        <p>
          The top 20 players at each nightly tournament will receive points
          based on their tournament placing. Players compete throughout the
          season to build up their points and ranking at each individual bar
          they play.
        </p>
        <p>
          Each season is approximately 18 weeks long. At the conclusion of the
          season the top 10 at each location qualify for an entry into our TOC
          Championship. Season point winners at each location will play a
          special FINAL championship table the night before TOC.
        </p>
        <p>
          At the conclusion of each season there will be a Season Ending Social
          (SES) that will be open to ALL PPC players who have played a minimum
          of 6 games at any of the host locations throughout the season.
        </p>
        <h2>Our Difference</h2>
        <hr />
        <p>
          Our games are fully managed by our well trained and courteous
          Tournament Directors. Our games are played on felt covered pine wood
          table-tops, with 100% plastic Copag® playing cards, provider of the
          cards used at the World Series of Poker.
        </p>
        <p>
          In addition, we offer more prizes to our players than any other league
          with simple, straight forward ways to win.
        </p>
      </b-col>
    </b-row>
  </b-container>
</template>
<script>
export default {
  computed: {
    atTop() {
      return this.$store.state.atTop
    },
    marginTop() {
      return this.$store.state.atTop ? 'marginTopTop' : 'marginTopOff'
    },
  },
}
</script>
<style>
.marginTopTop {
  margin-top: 180px;
}

.marginTopOff {
  margin-top: 130px;
}
</style>
