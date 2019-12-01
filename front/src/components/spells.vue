<template>
	<div>
		<h4>Spells</h4>
		<Spell
			v-for="(val, idx) in spells"
			v-bind:spell_id="val.id"
			:spellsList="spellsList"
			:key="val.id"
			v-on:remove="removeSpell(val, idx)"
			v-on:update="getSpells"
		></Spell>
		<div class="card">
			<h5 class="header" data-toggle="collapse" data-target="#new_spell" aria-expanded="false">
				New Spell
			</h5>
			<div class="collapse card-body p-1" id="new_spell">
				<Spell :spellsList="spellsList" @add="addSpell" @update="getSpells"></Spell>
			</div>
		</div>
		<hr />
	</div>
</template>

<script>
import Spell from './spell.vue'
export default {
	name: 'Spells',
	props: ['ref_id'],
	components: { Spell },
	watch: {
		ref_id: function(newVal) {
			this.get(newVal)
		},
	},
	created: function() {
		this.get(this.ref_id)
		this.getSpells()
	},
	data() {
		return {
			spellsList: [],
			spells: [],
		}
	},
	methods: {
		getSpells: function() {
			this.$http.get(process.env.VUE_APP_API_ENDPOINT + `/spells?lang=${this.$language}`).then(function(res) {
				console.debug(res)
				this.spellsList = res.data
			})
		},
		get: function(cardID) {
			this.$http
				.get(process.env.VUE_APP_API_ENDPOINT + `/ref/${cardID}/spell?lang=${this.$language}`)
				.then(function(res) {
					console.debug(res)
					this.spells = res.data
				})
		},
		removeSpell: function(spell, index) {
			this.$http
				.delete(process.env.VUE_APP_API_ENDPOINT + `/ref/${this.ref_id}/spell/${spell.id}`)
				.then(function(res) {
					console.log(res)
					if (res.status === 204) {
						this.spells.splice(index, 1)
					}
				})
		},
		addSpell: function(spell) {
			this.$http
				.put(process.env.VUE_APP_API_ENDPOINT + `/ref/${this.ref_id}/spell/${spell.id}?lang=${this.$language}`)
				.then(function(res) {
					console.log(res)
					if (res.status === 201) {
						this.spells.push(spell)
						this.spell = {}
					}
				})
		},
	},
}
</script>

<style lang="scss" scoped></style>
