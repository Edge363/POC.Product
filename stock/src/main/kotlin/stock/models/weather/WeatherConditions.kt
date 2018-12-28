package stock.models.weather

data class WeatherConditions(
        val cod: String? = "",
        val message: Double? = 0.0,
        val cnt: Int? = 0,
        val list: List<WeatherCondition?>? = null,
        val city: City? = null
)